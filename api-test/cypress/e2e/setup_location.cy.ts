import { faker } from "@faker-js/faker";

before(() => {
  cy.fixture("superadmin").then((superadmin) => {
    //
    cy.request({
      method: "POST",
      url: "/login",
      body: {
        username: superadmin.username,
        password: superadmin.password,
      },
    }).then((response) => {
      const accessToken = response.body.data.access_token;
      cy.wrap(accessToken).as("accessToken");
    });
  });

  cy.request({
    method: "POST",
    url: "/organizations/search",
    body: {
      name: "org1",
    },
  }).then((response) => {
    cy.wrap(response.body.data[0].id).as("OrganizationID");
  });
});

describe("Site", () => {
  beforeEach(function () {
    const accessToken = (this as any).accessToken as string;
    cy.setCookie("access_token", accessToken);
  });

  it("get all sites", () => {
    cy.request({
      method: "POST",
      url: "/sites/search",
      body: {
        all: true,
        with_organization: true,
        with_point: true,
      },
    });
  });

  it("creates site", function () {
    const organizationID = (this as any).OrganizationID as string;
    cy.request({
      method: "PUT",
      url: "/sites",
      body: {
        name: `site_${faker.number.int({ min: 1000, max: 9999 })}`,
        organization_id: organizationID,
      },
    }).then((response) => {
      cy.wrap(response.body.data.id).as("siteID");
    });
  });

  it("creates building", function () {
    const siteID = (this as any).siteID as string;
    cy.log(siteID);
    cy.request({
      method: "PUT",
      url: "/buildings",
      body: {
        name: `building_${faker.number.int({ min: 1000, max: 9999 })}`,
        external_name: `building_${faker.company.name()}`,
        site_id: siteID,
      },
    }).then((response) => {
      cy.wrap(response.body.data.id).as("buildingID");
    });
  });

  it("creates floor", function () {
    const buildingID = (this as any).buildingID as string;
    cy.log(buildingID);

    const floorNumber = faker.number.int({ min: 1, max: 15 });
    cy.request({
      method: "PUT",
      url: "/floors",
      body: {
        name: `floor_${floorNumber}`,
        number: floorNumber,
        building_id: buildingID,
      },
    }).then((response) => {
      cy.wrap(response.body.data.id).as("floorID");
    });
  });

  it("creates point", function () {
    const floorID = (this as any).floorID as string;
    cy.log(floorID);

    cy.request({
      method: "PUT",
      url: "/points",
      body: {
        name: `point_${faker.number.int({ min: 1000, max: 9999 })}`,
        external_name: `point_${faker.person.firstName()}`,
        floor_id: floorID,
      },
    }).then((response) => {
      cy.wrap(response.body.data.id).as("pointID");
    });
  });
});
