import { defineConfig } from "cypress";

export default defineConfig({
  projectId: 'jnqik5',
  e2e: {
    baseUrl: "http://localhost:8080/api/v1",
    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
  },
});
