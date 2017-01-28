/*
 *    Copyright (C) 2014 Daniel 'grindhold' Brendle
 *                  2014-2017 Christian Muehlhaeuser
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Daniel 'grindhold' Brendle <grindhold@skarphed.org>
 *      Christian Muehlhaeuser <muesli@gmail.com>
 */

package jenkinsbee

import (
	"github.com/muesli/beehive/bees"
)

type JenkinsBeeFactory struct {
	bees.BeeFactory
}

func (factory *JenkinsBeeFactory) New(name, description string, options bees.BeeOptions) bees.BeeInterface {
	bee := JenkinsBee{
		Bee:      bees.NewBee(name, factory.Name(), description, options),
		url:      options.GetValue("url").(string),
		user:     options.GetValue("user").(string),
		password: options.GetValue("password").(string),
	}
	return &bee
}

func (factory *JenkinsBeeFactory) Name() string {
	return "jenkinsbee"
}

func (factory *JenkinsBeeFactory) Description() string {
	return "A bee that triggers and reads info from Jenkins-Builds"
}

func (factory *JenkinsBeeFactory) Image() string {
	return factory.Name() + ".png"
}

func (factory *JenkinsBeeFactory) Options() []bees.BeeOptionDescriptor {
	opts := []bees.BeeOptionDescriptor{
		{
			Name:        "url",
			Description: "The url the jenkins-installation is reachable at",
			Type:        "string",
			Mandatory:   true,
		},
		{
			Name:        "user",
			Description: "HTTP auth username",
			Type:        "string",
		},
		{
			Name:        "password",
			Description: "HTTP auth password",
			Type:        "string",
		},
	}
	return opts
}

func (factory *JenkinsBeeFactory) Events() []bees.EventDescriptor {
	events := []bees.EventDescriptor{
		{
			Namespace:   factory.Name(),
			Name:        "statuschange",
			Description: "the status of a job has changed",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "name",
					Description: "Name of the job",
					Type:        "string",
				},
				{
					Name:        "status",
					Description: "Current status of the job ('red' or 'blue')",
					Type:        "string",
				},
				{
					Name:        "url",
					Description: "URL of the affected job",
					Type:        "string",
				},
			},
		},
	}
	return events
}

func (factory *JenkinsBeeFactory) Actions() []bees.ActionDescriptor {
	actions := []bees.ActionDescriptor{
		{
			Namespace:   factory.Name(),
			Name:        "trigger",
			Description: "Trigger a build on this jenkins machine",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "job",
					Description: "Name of the job on which to trigger a build",
					Type:        "string",
				},
			},
		},
	}
	return actions
}

func init() {
	f := JenkinsBeeFactory{}
	bees.RegisterFactory(&f)
}
