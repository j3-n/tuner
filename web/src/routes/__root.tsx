import {
  createRootRoute,
  Outlet
} from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'
import { LinkComponent } from '../components/link';

export const Route = createRootRoute({
  component: Page
});

function Page() {
  return (
    <>
      <LinkComponent to='/'></LinkComponent>
      <hr />
      <Outlet />
      {!import.meta.env.PROD && <TanStackRouterDevtools />}
    </>
  );
}
