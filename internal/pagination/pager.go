package pager

import (
	"context"

	metal "github.com/equinix-labs/metal-go/metal/v1"
)

func GetAllProjects(s metal.ProjectsApiService, include []string, exclude []string) ([]metal.Project, error) {
	var projects []metal.Project

	page := int32(1)     // int32 | Page to return (optional) (default to 1)
	perPage := int32(20) // int32 | Items returned per page (optional) (default to 10)
	for {
		projectPage, _, err := s.FindProjects(context.Background()).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
		if err != nil {
			return nil, err
		}
		projects = append(projects, projectPage.GetProjects()...)
		if projectPage.Meta.GetLastPage() > projectPage.Meta.GetCurrentPage() {
			page = page + 1
			continue
		}
		return projects, nil
	}
}

func GetDeviceEvents(s metal.ApiFindDeviceEventsRequest) ([]metal.Event, error) {
	var events []metal.Event

	page := int32(1)     // int32 | Page to return (optional) (default to 1)
	perPage := int32(20) // int32 | Items returned per page (optional) (default to 10)

	for {
		eventsPage, _, err := s.Page(page).PerPage(perPage).Execute()
		if err != nil {
			return nil, err
		}

		events = append(events, eventsPage.GetEvents()...)
		if eventsPage.Meta.GetLastPage() > eventsPage.Meta.GetCurrentPage() {
			page = page + 1
			continue
		}
		return events, nil
	}
}

func GetProjectEvents(s metal.ApiFindProjectEventsRequest) ([]metal.Event, error) {
	var events []metal.Event

	page := int32(1)     // int32 | Page to return (optional) (default to 1)
	perPage := int32(20) // int32 | Items returned per page (optional) (default to 10)

	for {
		eventsPage, _, err := s.Page(page).PerPage(perPage).Execute()
		if err != nil {
			return nil, err
		}

		events = append(events, eventsPage.GetEvents()...)
		if eventsPage.Meta.GetLastPage() > eventsPage.Meta.GetCurrentPage() {
			page = page + 1
			continue
		}
		return events, nil
	}
}

func GetOrganizationEvents(s metal.ApiFindOrganizationEventsRequest) ([]metal.Event, error) {
	var events []metal.Event

	page := int32(1)     // int32 | Page to return (optional) (default to 1)
	perPage := int32(20) // int32 | Items returned per page (optional) (default to 10)

	for {
		eventsPage, _, err := s.Page(page).PerPage(perPage).Execute()
		if err != nil {
			return nil, err
		}

		events = append(events, eventsPage.GetEvents()...)
		if eventsPage.Meta.GetLastPage() > eventsPage.Meta.GetCurrentPage() {
			page = page + 1
			continue
		}
		return events, nil
	}
}

func GetAllEvents(s metal.ApiFindEventsRequest) ([]metal.Event, error) {
	var events []metal.Event

	page := int32(1)     // int32 | Page to return (optional) (default to 1)
	perPage := int32(20) // int32 | Items returned per page (optional) (default to 10)
	for {
		eventsPage, _, err := s.Page(page).PerPage(perPage).Execute()
		if err != nil {
			return nil, err
		}

		events = append(events, eventsPage.GetEvents()...)
		if eventsPage.Meta.GetLastPage() > eventsPage.Meta.GetCurrentPage() {
			page = page + 1
			continue
		}
		return events, nil
	}
}

func GetAllOrganizations(s metal.OrganizationsApiService, include, exclude []string) ([]metal.Organization, error) {
	var orgs []metal.Organization
	page := int32(1)     // int32 | Page to return (optional) (default to 1)
	perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

	for {
		orgPage, _, err := s.FindOrganizations(context.Background()).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
		if err != nil {
			return nil, err
		}
		orgs = append(orgs, orgPage.GetOrganizations()...)
		if orgPage.Meta.GetLastPage() > orgPage.Meta.GetCurrentPage() {
			page = page + 1
			continue
		}
		return orgs, nil
	}
}

func GetProjectDevices(s metal.ApiFindProjectDevicesRequest) ([]metal.Device, error) {
	var devices []metal.Device

	page := int32(1)     // int32 | Page to return (optional) (default to 1)
	perPage := int32(20) // int32 | Items returned per page (optional) (default to 10)
	for {
		devicePage, _, err := s.Page(page).PerPage(perPage).Execute()
		if err != nil {
			return nil, err
		}
		devices = append(devices, devicePage.Devices...)
		if devicePage.Meta.GetLastPage() > devicePage.Meta.GetCurrentPage() {
			page = page + 1
			continue
		}
		return devices, nil
	}
}

func GetAllIPReservations(s metal.IPAddressesApiService, projectId string, inc []string, exc []string, types []metal.FindIPReservationsTypesParameterInner) ([]metal.IPReservationListIpAddressesInner, error) {
	var ipReservations []metal.IPReservationListIpAddressesInner
	page := int32(1)     // int32 | Page to return (optional) (default to 1)
	perPage := int32(20) // int32 | Items returned per page (optional) (default to 10)

	for {
		ipReservationsPage, _, err := s.FindIPReservations(context.Background(), projectId).Types(types).Include(inc).Exclude(exc).PerPage(perPage).Execute()
		if err != nil {
			return nil, err
		}

		ipReservations = append(ipReservations, ipReservationsPage.GetIpAddresses()...)
		if ipReservationsPage.Meta.GetLastPage() > ipReservationsPage.Meta.GetCurrentPage() {
			page = page + 1
			continue
		}
		return ipReservations, nil
	}
}

func GetAllProjectInterconnections(s metal.InterconnectionsApiService, projectId string, include []string, exclude []string) ([]metal.Interconnection, error) {
	var interConn []metal.Interconnection

	page := int32(1)     // int32 | Page to return (optional) (default to 1)
	perPage := int32(20) // int32 | Items returned per page (optional) (default to 10)

	for {
		interConnPage, _, err := s.ProjectListInterconnections(context.Background(), projectId).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
		if err != nil {
			return nil, err
		}
		interConn = append(interConn, interConnPage.GetInterconnections()...)
		if interConnPage.Meta.GetLastPage() > interConnPage.Meta.GetCurrentPage() {
			page = page + 1
			continue
		}
		return interConn, nil
	}
}
