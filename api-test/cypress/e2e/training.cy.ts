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

describe("Training", () => {
  it("start training", () => {
    cy.request({
      method: "PUT",
      url: "/training/",
      body: {
        training_name: "training1",
        training_type: "SUPERVISED",
      },
    });
  });
});
