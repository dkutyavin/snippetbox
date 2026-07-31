# What's that?

This is just a study project I assembled while following the book "Let's Go!" by
Alex Edwards, with some minor tweaks for my environment.

## How to start

I don't expect anyone but myself to run this app, but if you want to, here's how:

1. Install Go and Docker.
2. Add your secrets: `cp .env.example .env`. Feel free to change the values.
3. Prepare the environment: `make setup`.
4. Run the app: `make run`.

That's it! You can access the site via [https://localhost:4000](https://localhost:4000)
(note the secure protocol).
