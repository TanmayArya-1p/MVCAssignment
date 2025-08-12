
export default function NotFoundScreen() {
    return  <div className="min-h-screen min-w-screen flex flex-col items-center justify-center gap-10">
        <title>Not Found - InOrder</title>
        <a className="text-red-500"></a>
        <a className="text-yellow-500"></a>
        <a className="text-teal-500"></a>
        <a className="text-green-500"></a>
      <div className="ubuntu-bold text-7xl">Oops!</div>
      <div className="ubuntu-bold text-4xl">Page Not Found</div>
      <div className="ubuntu-bold text-xl">The page you are looking for does not exist.</div>
      <button className="flex flex-row gap-2" onClick={() => window.location.href="/"}>
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" className="size-6">
            <path strokeLinecap="round" strokeLinejoin="round" d="M10.5 19.5 3 12m0 0 7.5-7.5M3 12h18" />
            </svg>
        Go Back Home
      </button>
  </div>
}