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

describe("Files", () => {
  it("creates a bucket", () => {
    cy.request({
      method: "POST",
      url: "/files/bucket/bucket2",
    });
  });
});
