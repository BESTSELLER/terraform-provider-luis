package provider

import (
	"encoding/json"
	"strconv"

	"github.com/BESTSELLER/terraform-provider-luis/client"
	"github.com/BESTSELLER/terraform-provider-luis/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceClosedListSublist() *schema.Resource {
	return &schema.Resource{
		Create: resourceClosedListSublistCreate,
		Update: resourceClosedListSublistUpdate,
		Read:   resourceClosedListSublistRead,
		Delete: resourceClosedListSublistDelete,
		Importer: &schema.ResourceImporter{
			State: resourceClosedListSublistImport,
		},
		Description: "Configures a sublist to an existing closed list",

		Schema: map[string]*schema.Schema{
			"closed_list_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The closed list entity extractor ID.",
			},
			"canonical_form": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Canonical form is the unique identifier. When using a closed list entity in LUIS, there always is a canonical form and optional multiple synonyms",
			},
			"list": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required:    true,
				ForceNew:    false,
				Description: "List of synonyms",
			},
		},
	}
}

func resourceClosedListSublistCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.Client)
	listID := d.Get("closed_list_id").(string)
	var list []string

	for _, item := range d.Get("list").([]interface{}) {
		list = append(list, item.(string))
	}

	payload := models.ClosedListSublist{
		CanonicalForm: d.Get("canonical_form").(string),
		List:          list,
	}

	resp, err := c.SendRequest("POST", "closedlists/"+listID+"/sublists", payload, 201)
	if err != nil {
		return err
	}

	id := string(resp)
	d.SetId(id)
	return resourceClosedListSublistRead(d, m)
}

func resourceClosedListSublistUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.Client)
	listID := d.Get("closed_list_id").(string)
	var list []string

	for _, item := range d.Get("list").([]interface{}) {
		list = append(list, item.(string))
	}

	payload := models.ClosedListSublist{
		CanonicalForm: d.Get("canonical_form").(string),
		List:          list,
	}

	_, err := c.SendRequest("PUT", "closedlists/"+listID+"/sublists/"+d.Id(), payload, 200)
	if err != nil {
		return err
	}

	return resourceClosedListSublistRead(d, m)
}

func resourceClosedListSublistRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.Client)

	listID := d.Get("closed_list_id").(string)
	canonicalForm := d.Get("canonical_form").(string)

	body, err := c.SendRequest("GET", "closedlists/"+listID, nil, 200)
	if err != nil {
		return err
	}

	var closedListEntity models.ClosedListEntity
	err = json.Unmarshal(body, &closedListEntity)
	if err != nil {
		return err
	}

	id := 0
	if !d.IsNewResource() {
		id, err = strconv.Atoi(d.Id())
		if err != nil {
			return err
		}
	}

	found := false
	for _, entity := range closedListEntity.SubLists {

		// if new resource check for canonical, if existing check for ID
		if d.IsNewResource() {
			// check for canonical form
			if entity.CanonicalForm == canonicalForm {
				d.SetId(strconv.Itoa(entity.ID))
				d.Set("canonical_form", entity.CanonicalForm)
				d.Set("list", entity.List)
				found = true
				break
			}
		} else {
			// check for id
			if entity.ID == id {
				found = true
				d.Set("canonical_form", entity.CanonicalForm)
				d.Set("list", entity.List)
				break
			}
		}

	}

	if !found {
		d.SetId("")
	}

	return nil
}

func resourceClosedListSublistDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.Client)

	listID := d.Get("closed_list_id").(string)

	_, err := c.SendRequest("DELETE", "closedlists/"+listID+"/sublists/"+d.Id(), nil, 200)
	if err != nil {
		return err
	}

	return nil
}

func resourceClosedListSublistImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{d}, nil
}
