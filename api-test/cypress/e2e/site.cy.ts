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
});

describe("Site", () => {
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
});
