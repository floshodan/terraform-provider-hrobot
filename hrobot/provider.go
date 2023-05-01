package hrobot

import (
	"os"

	"github.com/floshodan/hrobot-go/hrobot"
	"github.com/floshodan/terraform-provider-hrobot/internal/boot"
	"github.com/floshodan/terraform-provider-hrobot/internal/firewall"
	"github.com/floshodan/terraform-provider-hrobot/internal/sshkey"
	"github.com/floshodan/terraform-provider-hrobot/internal/vswitch"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			//"hrobot_server_list": server.ResourceEvent(),
			//server.ResourceType: server.Resource(),
			sshkey.ResourceType:   sshkey.Resource(),
			firewall.ResourceType: firewall.Resource(),
			boot.ResourceType:     boot.Resource(),
			vswitch.ResourceType:  vswitch.Resource(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			sshkey.DataSourceType:      sshkey.DataSource(),
			sshkey.DataSourceListType:  sshkey.DataSourceList(),
			firewall.DataSourceType:    firewall.DataSource(),
			boot.DataSourceType:        boot.DataSourceList(),
			vswitch.DataSourceType:     vswitch.DataSource(),
			vswitch.DataSourceListType: vswitch.DataSourceList(),
			//rescue.DataSourceType:     rescue.DataSource(),
		},
		ConfigureFunc: providerConfigure,
	}
}
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	// TODO
	//username := d.Get("username").(string)
	//password := d.Get("password").(string)
	client := hrobot.NewClient(hrobot.WithToken(os.Getenv("HROBOT_TOKEN")))
	return client, nil
}
