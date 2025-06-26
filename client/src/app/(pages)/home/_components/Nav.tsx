import React from "react";
import Image from "next/image";
function Nav() {
  return (
    <div className=" items-center justify-center h-fit w-fit mt-4 ml-3 flex   gap-2 ">
      <Image
        alt="logo"
        src={"/logo.webp"}
        className=" size-8"
        height={100}
        width={100}
      />
      <span className="  font-semibold">ManiMate</span>
    </div>
  );
}

export default Nav;
