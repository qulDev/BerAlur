import { expect, test } from "@playwright/test";

test("web and API are reachable through one origin", async ({
  page,
  request,
}) => {
  await page.goto("/");
  await expect(
    page.getByRole("heading", { name: "Beralur, mulai dari fondasi." }),
  ).toBeVisible();
  await page.getByRole("link", { name: "Periksa API" }).click();
  await expect(page).toHaveURL(/\/health$/);
  const health = await request.get("/health");
  expect(health.status()).toBe(200);
  expect(await health.json()).toEqual({ status: "ok" });
  expect(health.headers()["x-request-id"]).toBeTruthy();
  const ready = await request.get("/ready");
  expect(ready.status()).toBe(200);
  expect(await ready.json()).toEqual({ status: "ready" });
  const missing = await request.get("/api/v1/missing");
  expect(missing.status()).toBe(404);
  expect(await missing.json()).toEqual({
    error: { code: "NOT_FOUND", message: "Resource not found" },
  });
});
