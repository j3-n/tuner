import {
  createRootRoute,
  Outlet
} from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'

export const Route = createRootRoute({
  component: Page
});

function Page() {
  return (
    <>
      <hr />
      <Outlet />
      {!import.meta.env.PROD && <TanStackRouterDevtools />}
    </>
  );
}
