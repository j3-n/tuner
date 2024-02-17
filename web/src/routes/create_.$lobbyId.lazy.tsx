import { createLazyFileRoute } from '@tanstack/react-router'

export const Route = createLazyFileRoute('/create/$lobbyId')({
  component: Page
});

function Page() {
  return (
    <></>
  )
}
