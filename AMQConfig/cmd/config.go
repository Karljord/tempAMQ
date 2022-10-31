package cmd

import (
	"fmt"
	"html/template"
	"os"

	"golang.org/x/exp/slices"
)

type Topic struct {
	Name  string   `yaml:"name"`
	Roles []string `yaml:"roles"`
}

type Queue struct {
	Name  string   `yaml:"name"`
	Roles []string `yaml:"roles"`
}

type Config struct {
	Users []struct {
		Name     string  `yaml:"name"`
		Password string  `yaml:"password"`
		Topics   []Topic `yaml:"topics"`
		Queues   []Queue `yaml:"queues"`
	} `yaml:"users"`
}

type authorizationConfig struct {
	Users  []User
	Groups []Group
	Queues []string
	Topics []string
}

type User struct {
	Name     string
	Password string
}

type Group struct {
	Name  string
	Users []string
}

func (config Config) generateAllData() (authConfig authorizationConfig) {

	tempGroups := make(map[string][]string)
	for _, user := range config.Users {

		authConfig.Users = append(authConfig.Users, User{user.Name, user.Password})

		for _, queue := range user.Queues {
			if !slices.Contains(authConfig.Queues, queue.Name) {
				authConfig.Queues = append(authConfig.Queues, queue.Name)
			}
			for _, role := range queue.Roles {
				topicGroupName := queue.Name + "." + role + "s"
				if !slices.Contains(tempGroups[topicGroupName], role+"s") {
					tempGroups[topicGroupName] = append(tempGroups[topicGroupName], user.Name)
				}
			}
		}

		for _, topic := range user.Topics {
			if !slices.Contains(authConfig.Topics, topic.Name) {
				authConfig.Topics = append(authConfig.Topics, topic.Name)
			}
			for _, role := range topic.Roles {

				topicGroupName := topic.Name + "." + role + "s"
				if !slices.Contains(tempGroups[topicGroupName], role+"s") {
					tempGroups[topicGroupName] = append(tempGroups[topicGroupName], user.Name)
				}
			}
		}
	}

	for groupName, users := range tempGroups {
		authConfig.Groups = append(authConfig.Groups, Group{groupName, users})
	}

	return
}

func (authConfig authorizationConfig) generateUserProperties() {
	for _, user := range authConfig.Users {
		fmt.Println(user.Name + "=" + user.Password)
	}
}

func (authConfig authorizationConfig) generateGroupProperties() {
	for _, group := range authConfig.Groups {
		fmt.Print(group.Name + "=")
		for i, user := range group.Users {
			if i == len(group.Users)-1 {
				fmt.Println(user)
			} else {
				fmt.Print(user + ",")
			}
		}
	}
}

func (authConfig authorizationConfig) generateAuthorizationEntries() {
	b, _ := os.ReadFile("/Users/leo/repo/rl/AMQConfig/cmd/template.xml")
	kebab := string(b)
	tmpl := template.Must(template.New("test").Parse(kebab))
	tmpl.Execute(os.Stdout, authConfig)
}
