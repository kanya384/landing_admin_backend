import { SectionsContext } from "./context/sectionsContext"
import { useBlocks } from "./hooks/sections.hook"


import { Loader } from './feature/loader';
import { useSections } from "./hooks/sections.hook";

function App() {
  const { blocks, setBlocks, menuClick, setMenuClick } = useSections()
  return (
    <div className="App container_main">
      <SectionsContext.Provider value={{ blocks, setBlocks, menuClick, setMenuClick }}>
        <div className="over">
          <Loader />
        </div>
      </SectionsContext.Provider>
    </div>
  );
}

export default App;