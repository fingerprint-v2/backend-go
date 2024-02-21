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
});
