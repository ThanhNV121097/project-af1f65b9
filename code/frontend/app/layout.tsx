import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "",
  description: "End-to-end greeting proof",
};

export default function RootLayout({ children }: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
