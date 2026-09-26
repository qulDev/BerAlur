import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "BerAlur — Fondasi awal",
  description: "Fondasi awal BerAlur untuk alur bisnis yang lebih tertata.",
};

export default function RootLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="id">
      <body>{children}</body>
    </html>
  );
}
