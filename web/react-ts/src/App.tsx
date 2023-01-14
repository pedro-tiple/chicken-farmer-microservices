import "./App.scss";
import { Farm } from "./views/Farm/Farm";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const queryClient = new QueryClient();

function App() {
  return (
    <div className="App">
      <QueryClientProvider client={queryClient}>
        <Farm/>
      </QueryClientProvider>
    </div>
  );
}

export default App;
