clear all;
close all;
clc;

system = 0;
betaLowerLimit = 25;
betaUpperLimit = 100;
betaStep = 1;

betas = betaLowerLimit:betaStep:betaUpperLimit;

n0 = 2000
n = 50
m = (200000 - n0)/n

AverageMagnetisations = [];
AverageEnergies = [];

figure;
hold on;
for beta = betas
    fname = ['./section2final2/beta-', num2str(beta) ,'-system', num2str(system) ,'.csv'];
    data = csvread(fname, 1);

    Sweep = data(:,1);
    Beta = data(:,2);
    SubSystemID = data(:,3);
    Magnetisation = data(:,4);
    DimensionlessEnergy = data(:,5);
        
    magnetisations = Magnetisation(n0:n:(n0+n*m));
    energies = DimensionlessEnergy(n0:n:(n0+n*m));
    AverageMagnetisations = [AverageMagnetisations mean(magnetisations)];
    AverageEnergies = [AverageEnergies mean(energies)];
end

plot(betas/100, AverageMagnetisations, 'x-')

hold off;
xlabel('$\beta$', 'Interpreter', 'latex', 'FontSize', 16);
ylabel('$<\mathcal{M}>$', 'Interpreter', 'latex', 'FontSize', 16);
% ylim([-1, 0.2])

% legend_handle = legend('-DynamicLegend');
% set(legend_handle,'Interpreter','latex')
% legend('show');

 
% figure;
% plot(Sweep, DimensionlessEnergy)
% ylabel('$E/J$', 'Interpreter', 'latex', 'FontSize', 16);