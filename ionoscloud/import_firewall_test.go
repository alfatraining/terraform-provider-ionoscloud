package ionoscloud

import (
	"fmt"

	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFirewall_ImportBasic(t *testing.T) {
	firewallName := "firewall"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckFirewallDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testacccheckfirewallconfigBasic, firewallName, firewallName),
			},
			{
				ResourceName:      "ionoscloud_firewall.webserver_http",
				ImportStateIdFunc: testAccFirewallImportStateId,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "ionoscloud_firewall.webserver_icmp",
				ImportStateIdFunc: testAccFirewallImportStateId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccFirewallImportStateId(s *terraform.State) (string, error) {
	var importID string = ""

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ionoscloud_firewall" {
			continue
		}

		importID = fmt.Sprintf("%s/%s/%s/%s", rs.Primary.Attributes["datacenter_id"], rs.Primary.Attributes["server_id"], rs.Primary.Attributes["nic_id"], rs.Primary.Attributes["id"])
	}

	return importID, nil
}
