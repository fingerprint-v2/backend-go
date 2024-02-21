import { genFingerprints } from "../../src/fake";

before(() => {
  cy.fixture("superadmin").then((superadmin) => {
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
    url: "/points/search",
    body: {
      name: "point1",
    },
  }).then((response) => {
    cy.wrap(response.body.data[0].id).as("pointId");
  });

  cy.request({
    method: "POST",
    url: "/sites/search",
    body: {
      name: "site1",
      with_organization: true,
    },
  }).then((response) => {
    cy.wrap(response.body.data[0].id).as("siteId");
  });
});

describe("perform surveys", () => {
  // Need to set cookie here or else cookie won't persist through all tests.
  beforeEach(function () {
    const accessToken = (this as any).accessToken as string;
    cy.setCookie("access_token", accessToken);
  });

  it("surveys supervisedly", function () {
    const pointId = (this as any).pointId as string;
    cy.log(pointId);
    const payload = {
      point_label_id: pointId,
      mode: "SUPERVISED",
      collect_device: {
        device_uid: "device3",
        device_id: "device_id3",
      },
      scan_mode: "INTERVAL",
      scan_interval: 1000,

      fingerprints: genFingerprints(10),
    };

    cy.request({
      method: "PUT",
      url: "/collect",
      body: payload,
    });
  });

  it("surveys unsupervisedly", function () {
    const siteID = (this as any).siteId as string;

    cy.log(siteID);

    const payload = {
      mode: "UNSUPERVISED",
      site_id: siteID,
      collect_device: {
        device_uid: "device3",
        device_id: "device_id3",
      },
      scan_mode: "INTERVAL",
      scan_interval: 1000,
      fingerprints: genFingerprints(10),
    };

    cy.request({
      method: "PUT",
      url: "/collect",
      body: payload,
    });
  });
});
