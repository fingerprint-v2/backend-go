describe("User", () => {
  it("passes", () => {
    const url = `http://localhost:8080/api/hello-world`;
    cy.request(url);
  });

  it("superadmin login", () => {
    //
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
        expect(response.body.data).to.have.property("access_token");
      });
    });
  });

  it("admin login", () => {
    //
    cy.fixture("admin").then((admin) => {
      //
      cy.request({
        method: "POST",
        url: "/login",
        body: {
          username: admin.username,
          password: admin.password,
        },
      }).then((response) => {
        expect(response.body.data).to.have.property("access_token");
      });
    });
  });

  it("user login", () => {
    //
    cy.fixture("user").then((user) => {
      //
      cy.request({
        method: "POST",
        url: "/login",
        body: {
          username: user.username,
          password: user.password,
        },
      }).then((response) => {
        expect(response.body.data).to.have.property("access_token");
      });
    });
  });
});
