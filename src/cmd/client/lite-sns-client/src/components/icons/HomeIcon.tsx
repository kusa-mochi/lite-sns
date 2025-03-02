import { css } from "@emotion/css"

type Props = {
    width?: string
    height?: string
    bodyColor?: string
    backgroundColor?: string
}

export default function HomeIcon(props: Props) {
    const { width, height, bodyColor, backgroundColor } = props

    const svgStyle = css`
        width: ${width ? width : "256px"};
        height: ${height ? height : "256px"};
    `
    const gStyle = css`
        background-color: ${backgroundColor ? backgroundColor : "#F0F0F0"};
    `
    const bodyStyle = css`
        fill: ${bodyColor ? bodyColor : "#4B4B4B"};
    `

    return (
        <svg className={svgStyle} version="1.1" id="_x32_" xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="0 0 512 512" xmlSpace="preserve">
            <g className={gStyle}>
                <path className={bodyStyle} d="M453.794,170.688L283.185,10.753c-15.287-14.337-39.083-14.337-54.37,0L58.206,170.688
                    c-8.012,7.515-12.565,18.01-12.565,29V472.25c0,21.954,17.803,39.75,39.75,39.75h120.947V395.145h99.324V512h120.946
                    c21.947,0,39.751-17.796,39.751-39.75V199.688C466.359,188.698,461.805,178.203,453.794,170.688z"></path>
            </g>
        </svg>
    )
}
