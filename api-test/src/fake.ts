import { faker } from "@faker-js/faker";

// const wifis = [
//   {
//     SSID: ".@ TrueMove H",
//     BSSID: "e4:4e:2d:8b:d0:8e",
//     capabilities: "[ESS]",
//     frequency: 5220,
//     level: -48,
//     timestamp: 1702449979,
//   },
//   {
//     SSID: ".@ TRUEWIFI",
//     BSSID: "e4:4e:2d:8b:d0:8f",
//     capabilities: "[ESS]",
//     frequency: 5220,
//     level: -48,
//     timestamp: 1702449979,
//   },
// ];

function genBSSID() {
  const BSSID = "XX:XX:XX:XX:XX:XX".replace(/X/g, function () {
    return "0123456789ABCDEF".charAt(Math.floor(Math.random() * 16));
  });
  return BSSID;
}

function genSSID() {
  return `wifi_${faker.person.firstName().toLocaleLowerCase()}`;
}

function genWifis(n: number) {
  const wifis = [];

  for (let i = 0; i < n; i++) {
    const wifi = {
      SSID: genSSID(),
      BSSID: genBSSID(),
      capabilities: "[ESS]",
      frequency: faker.number.int({ min: 2000, max: 6000 }),
      level: faker.number.int({ min: -100, max: -30 }),
      // Timestamp is not recorded in the database
      timestamp: faker.date.anytime().getTime(),
    };

    wifis.push(wifi);
  }

  return wifis;
}

export function genFingerprints(n: number) {
  const fingerprints = [];

  for (let i = 0; i < n; i++) {
    const fingerprint = {
      wifis: genWifis(faker.number.int({ min: 3, max: 10 })),
    };

    fingerprints.push(fingerprint);
  }

  return fingerprints;
}
