import { render, screen } from "@testing-library/react";
import { expect, test } from "vitest";
import Home from "./page";

test("provides same-origin links to check the API process and database readiness", () => {
  render(<Home />);
  expect(screen.getByRole("link", { name: "Periksa API" })).toHaveAttribute(
    "href",
    "/health",
  );
  expect(
    screen.getByRole("link", { name: "Periksa database" }),
  ).toHaveAttribute("href", "/ready");
});
