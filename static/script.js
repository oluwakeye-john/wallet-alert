(function () {
  // scoped

  const Ethereum = window.ethereum;

  let address = "";
  let currency = "";
  let email = "";

  const onMetamaskClick = async () => {
    console.log("--");
    if (Ethereum && Ethereum.isMetaMask) {
      const { error, account } = await getMetamaskAddress();

      if (error) {
        console.log(error);
      } else {
        console.log({ account });
        document.querySelector("#address-input").value = account;
      }
    } else {
      alert("Install metamask");
    }
  };

  const getMetamaskAddress = async () => {
    const request = {
      method: "eth_requestAccounts",
      params: [],
    };

    const defaultErrorMessage = "Error occurred";

    try {
      if (Ethereum.isConnected() && Ethereum.selectedAddress) {
        const account = Ethereum.selectedAddress;
        return { error: "", account };
      }

      const accounts = await Ethereum.request(request);

      if (accounts.length >= 1) {
        const account = accounts[0];
        return { error: "", account };
      } else {
        return defaultErrorMessage, "";
      }
    } catch (err) {
      return { error: err?.message || defaultErrorMessage, account: null };
    }
  };

  document
    .querySelector("#metamask-login-btn")
    .addEventListener("click", onMetamaskClick);

  window.addEventListener("load", () => {
    anime({
      targets: ".title, .subtitle, .form-1",
      translateY: [80, 0],
      opacity: [0, 1],
      easing: "spring",
      duration: 200,
      delay: anime.stagger(100),
    });
    document.querySelector(".form-1").style.pointerEvents = "all";
  });

  document.querySelector(".form-1").addEventListener("submit", (e) => {
    e.preventDefault();

    address = document.querySelector("#address-input").value;
    currency = document.querySelector("#currency-input").value;

    document.querySelector(".form-1").style.pointerEvents = "none";
    document.querySelector(".form-2").style.pointerEvents = "all";

    anime({
      targets: ".form-1",
      translateY: [0, 100],
      opacity: [1, 0],
      easing: "spring",
      duration: 200,
    });

    anime({
      targets: ".form-2",
      translateY: [100, 0],
      opacity: [0, 1],
      easing: "spring",
      duration: 200,
    });
  });

  document.querySelector(".form-2").addEventListener("submit", (e) => {
    e.preventDefault();
    email = document.querySelector("#email-input").value;
    console.log({ address, email, currency });
    alert("done");
  });
})();
