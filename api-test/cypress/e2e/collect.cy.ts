const wifis = [
  {
    SSID: ".@ TrueMove H",
    BSSID: "e4:4e:2d:8b:d0:8e",
    capabilities: "[ESS]",
    frequency: 5220,
    level: -48,
    timestamp: 1702449979,
  },
  {
    SSID: ".@ TRUEWIFI",
    BSSID: "e4:4e:2d:8b:d0:8f",
    capabilities: "[ESS]",
    frequency: 5220,
    level: -48,
    timestamp: 1702449979,
  },
];

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
      cy.setCookie("access_token", accessToken);
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

describe("Collect", () => {
  it("collects", function () {
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

      fingerprints: [
        {
          wifis: wifis,
        },
        {
          wifis: wifis,
        },
        {
          wifis: wifis,
        },
      ],
    };

    cy.request({
      method: "PUT",
      url: "/collect",
      body: payload,
    });
  });
});
