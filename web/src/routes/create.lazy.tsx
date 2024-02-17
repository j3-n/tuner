import { createLazyFileRoute } from '@tanstack/react-router'

export const Route = createLazyFileRoute('/create')({
  component: () => <div>Hello /create!</div>
})
