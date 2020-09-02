package schema

import (
	"amqp-subscriber/app/types"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func (c *Schema) Update(config types.SubscriberOption) (err error) {
	out, err := yaml.Marshal(config)
	if err != nil {
		return
	}
	if err = ioutil.WriteFile(
		c.autoload(config.Identity),
		out,
		0644,
	); err != nil {
		return
	}
	return
}