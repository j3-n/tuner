import { createLazyFileRoute } from '@tanstack/react-router'
import { ButtonComponent } from '../components/button';
import { InputComponent } from '../components/input';

export const Route = createLazyFileRoute('/lobby')({
  component: Page
});

function Page() {
  return (
    <div>
      <h1 className=''>Hello /lobby!</h1>

      <div className='flex'>
        <InputComponent placeholder='lobby code' />
        <ButtonComponent>hello</ButtonComponent>
      </div>
    </div>
  );
}
