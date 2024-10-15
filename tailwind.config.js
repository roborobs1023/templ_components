// const colors = require('tailwindcss/colors')

/** @type {import('tailwindcss').Config} */

module.exports = {
    content: [
        "./cmd/web/**/*.html", "./cmd/web/**/*.templ",
    ],
    theme: {
        // colors: {
        //     primary: colors.slate,
        //     white: colors.white,
        //     red: colors.red,
        //     slate: colors.slate,
        // },
        extend: {},
    },
    plugins: [
    ],
}

