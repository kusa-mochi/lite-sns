import {
  createContext,
  ReactNode,
  useContext,
  useEffect,
  useState,
} from "react";

type PaletteTheme = {
  main: string;
  fontColor: string;
  outlineColor: string;
};

type HeaderTheme = {
  fontSize: number;
};

type TitleTheme = {
  fontSize: number;
};

type BodyTheme = {
  fontSize: number;
};

type Theme = {
  palette: {
    primary: PaletteTheme;
    secondary: PaletteTheme;
    base: {
      bodyBackgroundColor: string;
      backgroundColor: string;
      cardColor: string;
      borderColor: string;
    };
  };
  typography: {
    fontFamily: string;
    fontSize: number;
    fontWeightLight: number;
    fontWeightRegular: number;
    fontWeightMedium: number;
    h1: HeaderTheme;
    h2: HeaderTheme;
    h3: HeaderTheme;
    h4: HeaderTheme;
    h5: HeaderTheme;
    h6: HeaderTheme;
    subTitle1: TitleTheme;
    subTitle2: TitleTheme;
    body1: BodyTheme;
    body2: BodyTheme;
    button: {
      textTransform: string;
    };
  };
};

type Themes = {
  light: Theme;
  dark: Theme;
};

const ThemeContext = createContext<Theme | null>(null);

// ThemeContextを使うコンポーネントが呼び出すフック
export function useTheme(): Theme {
  const theme = useContext(ThemeContext);
  if (!theme) throw new Error("wrap this component by ThemeProvider");

  return theme;
}

type ThemeProviderProps = {
  children: ReactNode;
  th: string;
};
export const ThemeProvider = (props: ThemeProviderProps) => {
  const [theme, setTheme] = useState<Theme | null>(null);
  const themesData: Themes = {
    light: {
      palette: {
        primary: {
          main: "#bbf7d0",
          fontColor: "rgba(0%, 0%, 0%, 0.87)",
          outlineColor: "#4abd9a",
        },
        secondary: {
          main: "#dde7ee",
          fontColor: "rgba(0%, 0%, 0%, 0.87)",
          outlineColor: "#4abd9a",
        },
        base: {
          backgroundColor: "#ffffff",
          bodyBackgroundColor: "#fafafc",
          cardColor: "#ffffff",
          borderColor: "#bbbbbb",
        },
      },
      typography: {
        fontFamily: "Arial, sans-serif",
        fontSize: 1, // "rem" size
        fontWeightLight: 300,
        fontWeightRegular: 400,
        fontWeightMedium: 700,

        h1: { fontSize: 60 },
        h2: { fontSize: 48 },
        h3: { fontSize: 42 },
        h4: { fontSize: 36 },
        h5: { fontSize: 20 },
        h6: { fontSize: 18 },
        subTitle1: { fontSize: 18 },
        subTitle2: { fontSize: 18 },
        body1: { fontSize: 16 },
        body2: { fontSize: 16 },
        button: { textTransform: "none" },
      },
    },
    dark: {
      // TODO: Make palette for the dark theme. Now it is same to the light theme.
      palette: {
        primary: {
          main: "#bbf7d0",
          fontColor: "rgba(0%, 0%, 0%, 0.87)",
          outlineColor: "#4abd9a",
        },
        secondary: {
          main: "#dde7ee",
          fontColor: "rgba(0%, 0%, 0%, 0.87)",
          outlineColor: "#4abd9a",
        },
        base: {
          backgroundColor: "#ffffff",
          bodyBackgroundColor: "#fafafc",
          cardColor: "#f0f1fa",
          borderColor: "#999999",
        },
      },
      typography: {
        fontFamily: "Arial, sans-serif",
        fontSize: 1, // "rem" size
        fontWeightLight: 300,
        fontWeightRegular: 400,
        fontWeightMedium: 700,

        h1: { fontSize: 60 },
        h2: { fontSize: 48 },
        h3: { fontSize: 42 },
        h4: { fontSize: 36 },
        h5: { fontSize: 20 },
        h6: { fontSize: 18 },
        subTitle1: { fontSize: 18 },
        subTitle2: { fontSize: 18 },
        body1: { fontSize: 16 },
        body2: { fontSize: 16 },
        button: { textTransform: "none" },
      },
    },
  };

  useEffect(() => {
    switch (props.th as keyof Themes) {
      case "light":
        setTheme(themesData.light);
        break;
      case "dark":
        setTheme(themesData.dark);
        break;
      default:
        throw new Error("invalid theme specification in provider");
    }
  }, []);

  if (!theme) return <div>Loading...</div>;

  return (
    <ThemeContext.Provider value={theme}>
      {props.children}
    </ThemeContext.Provider>
  );
};
