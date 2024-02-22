import { genFingerprints } from "../../src/fake";

before(function () {
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
    url: "/sites/search",
    body: {
      name: "site1",
      with_organization: true,
    },
  }).then((response) => {
    const siteID = response.body.data[0].id;
    cy.wrap(siteID).as("siteID");

    cy.request({
      method: "POST",
      url: "/points/search",
      body: {
        site_id: siteID,
      },
    }).then((response) => {
      const points = response.body.data;
      cy.wrap(points).as("points");
    });
  });
});

describe("perform surveys", () => {
  // Need to set cookie here or else cookie won't persist through all tests.
  beforeEach(function () {
    const accessToken = (this as any).accessToken as string;
    cy.setCookie("access_token", accessToken);
  });

  it("surveys supervisedly", function () {
    const points = (this as any).points as any[];
    cy.log(JSON.stringify(points));

    for (const point of points) {
      const payload = {
        point_label_id: point.id,
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
        url: "/localize/supervised",
        body: payload,
      });
    }
  });

  it("surveys unsupervisedly", function () {
    const siteID = (this as any).siteID as string;

    cy.log(siteID);

    const payload = {
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
      url: "/localize/unsupervised",
      body: payload,
    });
  });
});
