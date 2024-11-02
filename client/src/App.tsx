import TodoForm from "./components/TodoForm";
import TodoList from "./components/TodoList";

function App() {
  return (
    <>
      <div className="min-h-screen justify-center py-12 px-5 bg-gray-900">
        <div className="flex flex-col gap-y-5 max-w-2xl mx-auto">
          <TodoForm />
          <TodoList />
        </div>
      </div>
    </>
  );
}

export default App;
