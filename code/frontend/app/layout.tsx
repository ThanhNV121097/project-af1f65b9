import type { Metadata } from "next";
import "./globals.css";
import GreetingScreen from "../components/GreetingScreen";

export const metadata: Metadata = {
  title: "hello-word-2",
  description: "End-to-end greeting proof",
};

export default function RootLayout({ children }: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
