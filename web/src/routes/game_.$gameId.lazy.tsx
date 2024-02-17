import { createLazyFileRoute } from '@tanstack/react-router'

export const Route = createLazyFileRoute('/game/$gameId')({
  component: Page,
})

function Page() {
  const { gameId } = Route.useParams();

  return (
    <>welcome to {gameId}</>
  );
}
