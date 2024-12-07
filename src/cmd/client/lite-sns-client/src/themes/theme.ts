export default function useTheme(th: string) {
    const themes = {
        light: {
            palette: {
                primary: {
                    main: '#59d273',
                    fontColor: 'rgba(0%, 0%, 0%, 0.87)',
                },
                secondary: {
                    main: '#b3b3b3',
                    fontColor: 'rgba(0%, 0%, 0%, 0.87)',
                },
            },
            typography: {
                fontFamily: 'Arial, sans-serif',
                fontSize: 14,
                fontWeightLight: 300,
                fontWeightRegular: 400,
                fontWeightMedium: 700,
         
                h1: { fontSize: 60 },
                h2: { fontSize: 48 },
                h3: { fontSize: 42 },
                h4: { fontSize: 36 },
                h5: { fontSize: 20 },
                h6: { fontSize: 18 },
                subtitle1: { fontSize: 18 },
                subtitle2: { fontSize: 18 },
                body1: { fontSize: 16 },
                body2: { fontSize: 16 },
                button: { textTransform: 'none' },
            },
        },
        dark: { // TODO: Make palette for the dark theme. Now it is same to the light theme.
            palette: {
                primary: {
                    main: '#59d273',
                    fontColor: 'rgba(0%, 0%, 0%, 0.87)',
                },
                secondary: {
                    main: '##b3b3b3',
                    fontColor: 'rgba(0%, 0%, 0%, 0.87)',
                },
            },
            typography: {
                fontFamily: 'Arial, sans-serif',
                fontSize: 14,
                fontWeightLight: 300,
                fontWeightRegular: 400,
                fontWeightMedium: 700,
         
                h1: { fontSize: 60 },
                h2: { fontSize: 48 },
                h3: { fontSize: 42 },
                h4: { fontSize: 36 },
                h5: { fontSize: 20 },
                h6: { fontSize: 18 },
                subtitle1: { fontSize: 18 },
                subtitle2: { fontSize: 18 },
                body1: { fontSize: 16 },
                body2: { fontSize: 16 },
                button: { textTransform: 'none' },
            },
        },
    }

    switch (th) {
        case "light":
            return themes.light
        case "dark":
            return themes.dark
    }
}
