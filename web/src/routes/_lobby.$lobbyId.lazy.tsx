import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_lobby/$lobbyId')({
  component: Page
});

function Page() {
  return (
    <div>Hello /lobby/$lobbyId!</div>
  );
}
