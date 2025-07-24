import { useState, useCallback, useEffect } from "react";
import { Input } from "./components/ui/input";
import { ModeToggle } from "./components/theme-toggle";
import { collect } from "./service/engine";
import { useSearchParams } from "react-router-dom";


function App() {
  const [searchParams, setSearchParams] = useSearchParams();
  const [searchText, setSearchText] = useState<string>("");

  useEffect(() => {
    const queryFromURL = searchParams.get("q");
    if (queryFromURL) {
      setSearchText(queryFromURL);
    }
  }, []);

  const handleSearchSubmit = useCallback(async () => {
    if (searchText.trim() === "") {
      return;
    }
    setSearchParams({ q: searchText })

  }, [searchText]);

  const handleInputKeyPress = useCallback(
    (e: React.KeyboardEvent<HTMLInputElement>) => {
      if (e.key === "Enter") {
        handleSearchSubmit();
      }
    },
    [handleSearchSubmit],
  );

  useEffect(() => {
    const query = searchParams.get("q");

    const performSearch = async (q: string) => {
      if (q) {
        console.log("Initiating search for:", q);
        try {
          await collect(q);
          console.log("Search completed for:", q);
        } catch (err: any) {
          console.error("Error during search:", err);
        } finally {
        }
      }
    };

    if (query) performSearch(query);
    else setSearchText("");


  }, [searchParams]);

  return (
    <div className="flex w-full h-screen items-center justify-center">
      <div className="flex flex-row gap-2">
        <Input
          className="p-7"
          placeholder="Search..."
          value={searchText}
          onChange={(e) => {
            setSearchText(e.target.value);
          }}
          onKeyPress={handleInputKeyPress}
        />
        <ModeToggle className="absolute top-5 right-5" />
      </div>
    </div>
  );
}

export default App;
