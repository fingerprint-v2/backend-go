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
    url: "/sites/search",
    body: {
      name: "site1",
      with_organization: true,
    },
  }).then((response) => {
    cy.wrap(response.body.data[0].id).as("siteId");
  });
});

describe("Training", () => {
  it("start training", function () {
    const siteID = (this as any).siteId as string;
    cy.log(siteID);
    cy.request({
      method: "PUT",
      url: "/training/",
      body: {
        site_id: siteID,
        training_name: "training1",
        training_type: "SUPERVISED",
      },
    });
  });
});
