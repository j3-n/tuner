import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/lobby/$lobbyId')({
  component: Page
});

function Page() {
  const { lobbyId } = Route.useParams()

  return (
    <div>Hello /lobby/$lobbyId! {lobbyId}</div>
  );
}
