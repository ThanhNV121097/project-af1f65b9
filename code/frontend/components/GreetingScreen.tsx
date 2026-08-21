"use client";

import { useEffect, useState } from "react";
import styles from "./GreetingScreen.module.css";
import { greetingResponse, type GreetingState } from "../lib/mock/show-stored-greeting";

type ViewState = Exclude<GreetingState, "ready"> | "ready";

export default function GreetingScreen() {
  const [state, setState] = useState<ViewState>("loading");
  const [text, setText] = useState("");

  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    const forced = params.get("state") as ViewState | null;
    const nextState = forced ?? "ready";

    const timer = window.setTimeout(() => {
      if (nextState === "error") {
        setState("error");
        return;
      }
      if (nextState === "empty") {
        setState("empty");
        return;
      }
      setText(greetingResponse.text);
      setState("ready");
    }, 250);

    return () => window.clearTimeout(timer);
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
