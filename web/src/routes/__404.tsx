import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/__404')({
  component: Page,
  errorComponent: Page,
  pendingComponent: Page
});

function Page() {
    return (
      <div className="grid h-screen place-content-center bg-white px-4">
        <h1 className="">404 | Not Found</h1>
      </div>
    );
  }