export interface TestSquareProps {
  className: string;
}

function TestSquare({ className }: TestSquareProps) {
  return (
  <div>
    <div className={"w-[40px] h-[40px] text-black-400 bg-white-400 " + className}></div>
  </div>
  );
}

export default TestSquare;
