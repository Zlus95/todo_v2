import React from "react";

const Auth = () => (
  <div className="flex justify-center items-center h-full">
    <div className="h-80 w-72 bg-black/50 border-2 border-neutral-400 rounded">
      <div className="text-primary flex justify-center pt-4 text-lg">
        Todo List
      </div>
      <div className="flex flex-col gap-2 p-4">
        <label htmlFor="email" className="text-white/50">
          Email:
        </label>
        <input />
        <label htmlFor="password" className="text-white/50">
          Password:
        </label>
        <input />
      </div>
      <div className="flex justify-center pt-4">
        <button className="text-primary">Sign in</button>
      </div>
      <div className="p-4 flex gap-2">
        <p className="text-white/50">Don't have an account?</p>
        <p className="text-primary">Sign up</p>
      </div>
    </div>
  </div>
);

export default Auth;
