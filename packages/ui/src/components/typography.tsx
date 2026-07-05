import { cn } from "../lib/utilities";

const TypographyH1 = ({ className, ...props }: React.ComponentProps<"h1">) => {
  return (
    <h1
      className={cn(
        "scroll-m-20 text-center text-4xl font-extrabold tracking-tight text-balance",
        className
      )}
      {...props}
    >
      Taxing Laughter: The Joke Tax Chronicles
    </h1>
  );
};

const TypographyH2 = ({ className, ...props }: React.ComponentProps<"h2">) => {
  return (
    <h2
      className={cn(
        "scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight first:mt-0",
        className
      )}
      {...props}
    >
      The People of the Kingdom
    </h2>
  );
};

const TypographyH3 = ({ className, ...props }: React.ComponentProps<"h3">) => {
  return (
    <h3
      className={cn(
        "scroll-m-20 text-2xl font-semibold tracking-tight",
        className
      )}
      {...props}
    >
      The Joke Tax
    </h3>
  );
};

const TypographyH4 = ({ className, ...props }: React.ComponentProps<"h4">) => {
  return (
    <h4
      className={cn(
        "scroll-m-20 text-xl font-semibold tracking-tight",
        className
      )}
      {...props}
    >
      People stopped telling jokes
    </h4>
  );
};

const TypographyP = ({ className, ...props }: React.ComponentProps<"p">) => {
  return (
    <p
      className={cn("leading-7 [&:not(:first-child)]:mt-6", className)}
      {...props}
    >
      The king, seeing how much happier his subjects were, realized the error of
      his ways and repealed the joke tax.
    </p>
  );
};

const TypographyBlockquote = ({
  className,
  ...props
}: React.ComponentProps<"blockquote">) => {
  return (
    <blockquote
      className={cn("mt-6 border-l-2 pl-6 italic", className)}
      {...props}
    >
      &quot;After all,&quot; he said, &quot;everyone enjoys a good joke, so
      it&apos;s only fair that they should pay for the privilege.&quot;
    </blockquote>
  );
};

const TypographyList = ({
  className,
  ...props
}: React.ComponentProps<"ul">) => {
  return (
    <ul className={cn("my-6 ml-6 list-disc [&>li]:mt-2", className)} {...props}>
      <li>1st level of puns: 5 gold coins</li>
      <li>2nd level of jokes: 10 gold coins</li>
      <li>3rd level of one-liners : 20 gold coins</li>
    </ul>
  );
};

const TypographyInlineCode = ({
  className,
  ...props
}: React.ComponentProps<"code">) => {
  return (
    <code
      className={cn(
        "bg-muted relative rounded px-[0.3rem] py-[0.2rem] font-mono text-sm font-semibold",
        className
      )}
      {...props}
    >
      @radix-ui/react-alert-dialog
    </code>
  );
};

const TypographyLead = ({ className, ...props }: React.ComponentProps<"p">) => {
  return (
    <p className={cn("text-muted-foreground text-xl", className)} {...props}>
      A modal dialog that interrupts the user with important content and expects
      a response.
    </p>
  );
};

const TypographyLarge = ({
  className,
  ...props
}: React.ComponentProps<"div">) => {
  return (
    <div className={cn("text-lg font-semibold", className)} {...props}>
      Are you absolutely sure?
    </div>
  );
};

const TypographySmall = ({
  className,
  ...props
}: React.ComponentProps<"small">) => {
  return (
    <small
      className={cn("text-sm leading-none font-medium", className)}
      {...props}
    >
      Email address
    </small>
  );
};

const TypographyMuted = ({
  className,
  ...props
}: React.ComponentProps<"p">) => {
  return (
    <p className={cn("text-muted-foreground text-sm", className)} {...props}>
      Enter your email address.
    </p>
  );
};

export {
  TypographyH1,
  TypographyH2,
  TypographyH3,
  TypographyH4,
  TypographyP,
  TypographyBlockquote,
  TypographyList,
  TypographyInlineCode,
  TypographyLead,
  TypographyLarge,
  TypographySmall,
  TypographyMuted,
};
