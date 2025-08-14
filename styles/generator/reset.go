package generator

var resetCSS = `:root {
    interpolate-size: allow-keywords;
    transition-behavior: allow-discrete;
}

*,
:before,
:after {
    box-sizing: border-box;
    margin: 0;
    padding: 0;
}

h1,
h2,
h3,
h4,
h5,
h6 {
    text-wrap: balance;
    line-height: 1.2;
}

p,
li,
figcaption {
    text-wrap: balance;
}

b,
strong {
    font-weight: bolder;
}

/* body> :is(header, footer),
main,
section,
article {
    container-type: inline-size;
} */

ol,
ul,
menu {
    list-style: none;
    margin: 0;
    padding: 0;
}


a {
    text-decoration: inherit;
    color: inherit;

}

input,
button,
textarea,
select {
    font: inherit;
}

p,
h1,
h2,
h3,
h4,
h5,
h6 {
    overflow-wrap: break-word;
}

p {
    max-width: 75ch;
    text-wrap: stable;
}

h1,
h2,
h3,
h4,
h5,
h6 {
    text-wrap: balance;
}

#root,
#__next {
    isolation: isolate;
}

textarea {
    resize: vertical;
}

img,
picture,
video,
canvas,
svg {
    max-width: 100%;
    height: auto;
    vertical-align: middle;
    font-style: italic;
    font-size: 0.75rem;
    background-repeat: no-repeat;
    background-size: cover;
    shape-margin: 1rem;
}


body {
    min-height: 100vh;
    min-height: 100dvh;
    min-block-size: 100vh;
    min-block-size: 100dvh;
}

button,
input:where([type='button']),
input:where([type='reset']),
input:where([type='submit']) {
    -webkit-appearance: button;
    appearance: auto;
    /* 1 */
    background-color: transparent;
    /* 2 */
    background-image: none;
    /* 2 */
}

/*
Set the default cursor for buttons.
*/

button,
[role="button"] {
    cursor: pointer;
}
`

func (cg *CSSBuilder) GenerateResetCSS() {
	cg.Output.WriteString(resetCSS)
}
