package hubspot

import (
	"context"
//	"fmt"

	"github.com/tnelson-doghouse/hubspot"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

//// TABLE DEFINITION

func tableCompany(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hubspot_company",
		Description: "Company in Hubspot",
		List: &plugin.ListConfig{
			Hydrate: listCompanies,
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
//				Func:           getUserGroups,
				// Jira limited concurrency to avoid a 429 too many requests error, so I decided it'd be polite to do the same
				// Additionally, see https://developers.hubspot.com/docs/api/usage-details#rate-limits for rate limits
				MaxConcurrency: 50,
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the company.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "domain",
				Description: "The domain",
				Type:        proto.ColumnType_STRING,
//				Transform:   transform.FromGo(),
			},
			{
				Name:        "industry",
				Description: "The industry that the company works in",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "city",
				Description: "The city the company is located in",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "state",
				Description: "The state the company is located in",
				Type:        proto.ColumnType_STRING,
//				Transform:   transform.FromField("Active"),
			},
			{
				Name:        "phone",
				Description: "The phone number of the company",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "createddate",
				Description: "The date the company was added to Hubspot",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "lastmodifieddate",
				Description: "The date the company was last modified in HubSpot",
				Type:        proto.ColumnType_STRING,
//				Hydrate:     getUserGroups,
//				Transform:   transform.From(groupNames),
			},

			// Standard columns
			{
				Name:        "title",
				Description: "Company Name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func listCompanies(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_company.listCompanies", "connection_error", err)
		return nil, err
	}

	// If the requested number of items is less than the paging max limit
	// set the limit to that instead
/*	queryLimit := d.QueryContext.Limit
	var maxResults int = 1000
	if d.QueryContext.Limit != nil {
		if *queryLimit < 1000 {
			maxResults = int(*queryLimit)
		}
	} */

	nextlink := ""
	var nextnil hubspot.PagingResponsePage
	for {
//		apiEndpoint := fmt.Sprintf("rest/api/2/users/search?startAt=%d&maxResults=%d", last, maxResults)
//		apiEndpoint := fmt.Sprintf("/crm/v3/objects/companies")

		companies, err := client.CompaniesList(nextlink);

		if err != nil {
			plugin.Logger(ctx).Error("hubspot_company.listCompanies", "get_request_error", err)
			return nil, err
		}

/*		companies := new([]hubspot.Company)
		_, err = client.Do(req, companies)
		if err != nil {
			plugin.Logger(ctx).Error("hubspot_company.listCompanies", "api_error", err)
			return nil, err
		} */

		for _, company := range companies.Results {
			d.StreamListItem(ctx, company)
			// Context may get cancelled due to manual cancellation or if the limit has been reached
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
/*
		// evaluate paging start value for next iteration
		last = last + len(*companies)

		// API doesn't gives paging parameters in the response,
		// therefore using output length to quit paging
		if len(*companies) < 1000 {
			return nil, nil
		}

                last = resp.StartAt + len(issues)
		last = resp.paging
                if last >= resp.Total {
                        return nil, nil
                } else {
                        options.StartAt = last
                } */
		if companies.Paging.Next != nextnil {
			nextlink = companies.Paging.Next.Link
		} else {
			return nil, nil
		}
	}
}

//// HYDRATE FUNCTIONS
/*
func getUserGroups(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	user := h.Item.(jira.User)

	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_company.getUserGroups", "connection_error", err)
		return nil, err
	}

	groups, _, err := client.User.GetGroups(user.AccountID)
	if err != nil {
		plugin.Logger(ctx).Error("hubspot_company.getUserGroups", "api_error", err)
		return nil, err
	}

	return groups, nil
}

//// TRANSFORM FUNCTION

func groupNames(_ context.Context, d *transform.TransformData) (interface{}, error) {
	userGroups := d.HydrateItem.(*[]jira.UserGroup)
	var groupNames []string
	for _, group := range *userGroups {
		groupNames = append(groupNames, group.Name)
	}
	return groupNames, nil
}

*/
