"use client";

import { useEffect, useState } from "react";
import styles from "./GreetingScreen.module.css";

type GreetingResponse = { text: string };
type GreetingState = "ready" | "loading" | "error" | "empty";
type ViewState = Exclude<GreetingState, "ready"> | "ready";

const apiBase = process.env.NEXT_PUBLIC_API_URL ?? "/api";

export default function GreetingScreen() {
  const [state, setState] = useState<ViewState>("loading");
  const [text, setText] = useState("");

  useEffect(() => {
    const controller = new AbortController();

    fetch(`${apiBase}/v1/greeting`, { signal: controller.signal })
      .then(async (response) => {
        if (!response.ok) {
          if (response.status === 422) {
            setState("empty");
            return;
          }
          setState("error");
          return;
        }

        const data = (await response.json()) as GreetingResponse;
        if (data.text === "") {
          setState("empty");
          return;
        }

        setText(data.text);
        setState("ready");
      })
      .catch(() => {
        if (!controller.signal.aborted) {
          setState("error");
        }
      });

    return () => controller.abort();
  }, []);

  return (
    <main className={styles.frame}>
      <section className={styles.screen} aria-live="polite" aria-label="Greeting screen">
        {state === "ready" ? (
          <h1 className={styles.greeting}>{text}</h1>
        ) : (
          <p className={styles.message}>{state === "loading" ? "Loading greeting…" : state === "empty" ? "Greeting is empty." : "Greeting is not available."}</p>
        )}
      </section>
    </main>
  );
}
