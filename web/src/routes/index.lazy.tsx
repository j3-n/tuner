import { createLazyFileRoute } from '@tanstack/react-router';
import { LinkComponent } from '../components/link';

export const Route = createLazyFileRoute('/')({
  component: Page,
})

function Page() {
  return (
    <div className="p-2 fixed justify-center items-center">
      <h1>welcome to tuner!</h1>

      <LinkComponent to='/create'>create game</LinkComponent>
      <LinkComponent to='/lobby'>join game</LinkComponent>
    </div>
  );
}
