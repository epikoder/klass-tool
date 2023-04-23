/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"reflect"
	"strings"
	"time"

	"github.com/epikoder/klass-tool/src/database"
	"github.com/epikoder/klass-tool/src/models"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var (
	db          = database.DBold()
	db_         = database.DBnew()
	maxQuestion = 150
)

type (
	Question struct {
		models.Question
		TopicId   uint
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	Topic struct {
		Id        uint       `sql:"primary_key;" json:"id"`
		SubjectId int        `json:"subject_id"`
		Title     string     `json:"title"`
		Details   string     `json:"details"`
		Questions []Question `json:"questions"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	Subject struct {
		Id          uint   `sql:"primary_key;" json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Topics      []Topic
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db_.Migrator().DropTable(&Subject{})
		db_.Migrator().DropTable(&models.Source{})
		db_.Migrator().DropTable(&Topic{})
		db_.Migrator().DropTable(&Question{})
		db_.Migrator().DropTable(&models.User{})
		db_.AutoMigrate(&Subject{})
		db_.AutoMigrate(&Topic{})
		db_.AutoMigrate(&Question{})
		db_.AutoMigrate(&models.User{})

		data := getRecords()
		sources := getSource()
		sourceIds := []int{}
		for _, s := range sources {
			if s.Level == "WAEC" || s.Level == "JAMB" {
				sourceIds = append(sourceIds, int(s.Id))
			}
		}

		for _, subject := range data {
			if len(subject.Topics) == 0 {
				continue
			}
			s := &Subject{
				Name:        subject.Name,
				Description: subject.Description,
			}

			for _, topic := range subject.Topics {
				t := Topic{
					Title:   topic.Title,
					Details: topic.Details,
				}

				fnGetQuestions := func() []Question {
					ids := []int{}
					for _, o := range subject.Objectives {
						for _, oq := range o.Questions {
							for _, ot := range topic.Objectives {
								if ot.ObjectiveId == int(o.Id) {
									ids = append(ids, oq.QuestionId)
								}
							}
						}
					}
					ids = ids[:func() int {
						if len(ids) > maxQuestion {
							return maxQuestion
						}
						return len(ids)
					}()]
					questions := []Question{}
					if err := db.Find(&questions, "source_id IN ? AND id IN ? ORDER BY RAND()", sourceIds, ids).Error; err != nil {
						panic(err)
					}

					for index := range questions {
						questions[index].Id = nil
					}
					return questions
				}
				questions := fnGetQuestions()
				if len(questions) < 50 {
					continue
				}
				t.Questions = questions
				s.Topics = append(s.Topics, t)
			}
			if err := db_.Create(s).Error; err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func getRecords() []models.Subject {
	subjects := []models.Subject{}
	if err := db.
		Preload("Topics").
		Preload("Topics.Objectives").
		Preload("Objectives").
		Preload("Objectives.Questions").
		Find(&subjects).
		Error; err != nil {
		panic(err)
	}
	ss := []models.Subject{}
	for _, s := range subjects {
		ok := func() bool {
			for _, s_ := range ss {
				if strings.EqualFold(s.Name, s_.Name) {
					return false
				}
			}
			return true
		}()
		if ok {
			ss = append(ss, s)
		}
	}
	return ss
}

func getSource() []models.Source {
	sources := []models.Source{}
	if err := db.Find(&sources, "level = ? OR level = ? ORDER BY RAND()", "WAEC", "JAMB").Error; err != nil {
		panic(err)
	}
	return sources
}

func insertRecords(i interface{}, model interface{}) (errr error) {
	c := chunkSlice(i, 500)
	for _, d := range c {
		if err := db_.Model(model).Create(d).Error; err != nil {
			return
		}
	}
	return
}

func chunkSlice(slice interface{}, chunkSize int) []interface{} {
	var chunks []interface{}
	v := reflect.ValueOf(slice)
	switch v.Type().Kind() {
	case reflect.Slice:
		{
			for i := 0; i < v.Len(); i += chunkSize {
				end := i + chunkSize
				if end > v.Len() {
					end = v.Len()
				}
				chunks = append(chunks, v.Slice(i, end).Interface())
			}

			return chunks
		}
	default:
		{
			return chunks
		}
	}
}
