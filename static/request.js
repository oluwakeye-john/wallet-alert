const createSubscription = async (data) => {
  const { email, address, currency } = data;

  const query = `
    mutation {
        createSubscription(input: { address: "${address}", email: "${email}", currency_code: ${currency} }) {
            address
            is_subscribed
        }
    }
  `;

  return fetch("/query", {
    body: JSON.stringify({ query }),
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((res) => res.json())
    .then((res) => {
      if (res.errors) {
        console.log(res.errors, res.errors[0].message);
        return res.errors[0].message;
      }
      return "Done";
    })
    .catch((err) => {
      console.log({ err });
      return "An error occurred";
    });
};
