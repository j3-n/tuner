import { createLazyFileRoute } from '@tanstack/react-router';
import { LinkComponent } from '../components/link';
import Cookies from 'js-cookie';

export const Route = createLazyFileRoute('/')({
  component: Page,
})

function Page() {
  return (
    <div className="p-2 fixed justify-center items-center">
      <h1>welcome to tuner!</h1>

      <p>{Cookies.get("TUNER_SESSION")}</p>

      <LinkComponent to='/create'>Create Game</LinkComponent>
      <LinkComponent to='/lobby'>Join Game</LinkComponent>
      <LinkComponent to='/lobby'>Login with Spotify</LinkComponent>
    </div>
  );
}
