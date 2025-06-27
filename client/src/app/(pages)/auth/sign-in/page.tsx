"use client";
import { handleSignIn } from "@/handler/user-apis";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import Image from "next/image";
import Link from "next/link";
import { useRouter } from "next/navigation";
import React, { useEffect, useState } from "react";
import { toast } from "sonner";

function SignIn() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [disableSubmit, setDisableSubmit] = useState(true);

  const router = useRouter();
  useEffect(() => {
    if (password && email) {
      setDisableSubmit(false);
    } else {
      setDisableSubmit(true);
    }
  }, [email, password]);

  const handleSubmit = async (e: React.SyntheticEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      setDisableSubmit(true);
      const res = await handleSignIn(email, password);
      toast.success(res.message);
      router.push("/home");
    } catch (error: any) {
      toast.error(error.message);
    } finally {
      setDisableSubmit(false);
    }
  };
  return (
    <div className=" ">
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

      <div className=" h-full  w-full flex items-center justify-center ">
        <div className=" h-full w-[24rem]   px-2">
          <div className=" flex flex-col gap-4  items-center justify-center">
            <div>
              <Image
                alt="logo"
                src={"/logo.webp"}
                className=" size-20"
                height={400}
                width={400}
              />
            </div>
            <div className="   gap-2 flex flex-col items-center justify-center">
              <h1 className=" font-bold text-2xl">Sign in to ManiMate</h1>
              <div className=" flex flex-col  items-center justify-center text-xs text-neutral-400">
                <p>Get 1000 credits + 300 daily credits</p>
                <p>No waitlist—start creating now</p>
              </div>
            </div>
          </div>
          <div className=" flex flex-col gap-3 w-full mt-4 ">
            <div>
              <Button className=" bg-sidebar-accent hover:bg-sidebar-accent/70 w-full">
                <span>
                  <Image
                    src={"/google.svg"}
                    width={100}
                    height={100}
                    className=" size-4"
                    alt="google"
                  />
                </span>
                Sign in with Google
              </Button>
            </div>
            <div>
              <Button className=" bg-sidebar-accent hover:bg-sidebar-accent/70 w-full">
                <span>
                  <Image
                    src={"/github.png"}
                    width={100}
                    height={100}
                    className=" size-4"
                    alt="google"
                  />
                </span>
                Sign in with Github
              </Button>
            </div>
          </div>

          <div className="flex items-center gap-2 my-4 w-full">
            <div className="flex-1 h-px bg-neutral-800" />
            <span className="text-neutral-600 text-sm font-medium px-2">
              or
            </span>
            <div className="flex-1 h-px bg-neutral-800" />
          </div>

          <form onSubmit={handleSubmit} className=" flex flex-col gap-4">
            <div className=" flex flex-col gap-2">
              <label className=" text-xs">Email</label>
              <Input
                className=" rounded-lg"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
            </div>

            <div className=" flex flex-col gap-2">
              <label className=" text-xs">Password</label>
              <Input
                className=" rounded-lg"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>
            <Button type="submit" disabled={disableSubmit}>
              Sign in
            </Button>
          </form>

          <div className=" text-xs text-neutral-500  flex  items-center justify-center mt-5 gap-2">
            Don&apos;t have an account?{" "}
            <Link href={"/auth/sign-up"} className=" underline">
              Sign-up
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}

export default SignIn;
