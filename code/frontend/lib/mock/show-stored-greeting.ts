export type GreetingResponse = { text: string };

export type GreetingState = "ready" | "loading" | "error" | "empty";

export const greetingResponse: GreetingResponse = { text: "Hello Word" };
