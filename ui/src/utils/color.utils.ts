

export function randomRGBAColorGenerator(opacity: number) {
    return `rgba(${Math.floor(Math.random() * 255)},${Math.floor(Math.random() * 255)},${Math.floor(Math.random() * 255)},${opacity})`;
}